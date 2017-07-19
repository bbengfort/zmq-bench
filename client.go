package zmqnet

import (
	"time"

	pb "github.com/bbengfort/zmqnet/msg"
	"github.com/gogo/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
)

//===========================================================================
// Client for Remote Peers
//===========================================================================

// Client communicates with a remote peer.
type Client struct {
	Transporter
	messages uint64        // number of messages sent to measure throughput
	latency  time.Duration // total time to send messages for throughput
}

// Connect to the remote peer
func (c *Client) Connect() (err error) {
	// Create the socket
	if c.sock, err = c.net.context.NewSocket(zmq.REQ); err != nil {
		return err
	}

	// Connect to the server
	ep := c.host.ZMQEndpoint(false)
	if err = c.sock.Connect(ep); err != nil {
		return err
	}
	info("connected to %s\n", ep)

	return nil
}

// Reset the socket by setting the linger to 0, closing it, then reconnecting.
func (c *Client) Reset() error {

	// Close the socket
	if err := c.Close(); err != nil {
		return err
	}

	// And reconnect
	return c.Connect()
}

//===========================================================================
// Transport Methods
//===========================================================================

// Send a message to the remote peer in a safe fashion, specifying the # of
// retries and the timeout to wait on.
func (c *Client) Send(message string, mtype pb.MessageType, retries int, timeout time.Duration) error {
	if err := c.send(message, mtype); err != nil {
		return err
	}

	for {

		// Poll socket for a reply, with timeout
		poller := zmq.NewPoller()
		poller.Add(c.sock, zmq.POLLIN)
		sockets, err := poller.PollAll(timeout)
		if err != nil {
			return err
		}

		// Process a reply and exit if the reply is valid. Otherwise clsoe
		// socket and retry the message for num retries. Abandon after we
		// exhaust the number of allocated retries.
		if sock := sockets[0]; sock.Events&zmq.POLLIN != 0 {
			data, err := sock.Socket.RecvBytes(0)
			if err != nil {
				return err
			}

			reply := new(pb.Message)
			if err := proto.Unmarshal(data, reply); err != nil {
				return err
			}

			info("received: %s\n", reply.String())
			return nil

		} else if retries--; retries == 0 {
			warn("connection to %s is offline, message dropped", c.host.ZMQEndpoint(false))
			if err := c.Reset(); err != nil {
				return err
			}
			return nil
		} else {
			warn("no response from server, retrying send")

			// Old socket is confused, reset it.
			if err := c.Reset(); err != nil {
				return err
			}

			// Resend the original message
			if err := c.send(message, mtype); err != nil {
				return err
			}
		}
	}
}
