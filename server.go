package zmqnet

import (
	"fmt"

	pb "github.com/bbengfort/zmqnet/msg"
	zmq "github.com/pebbe/zmq4"
)

//===========================================================================
// Server and HTTP Transport
//===========================================================================

// Server responds to requests from other peers.
type Server struct {
	Transporter
}

// Run the server and listen for messages
func (s *Server) Run() (err error) {

	// Create the socket
	if s.sock, err = s.net.context.NewSocket(zmq.REP); err != nil {
		return err
	}

	// Bind the socket and run the listener
	ep := s.host.ZMQEndpoint(true)
	if err := s.sock.Bind(ep); err != nil {
		return err
	}
	info("bound to %s\n", ep)

	for {
		msg, err := s.recv()
		if err != nil {
			warne(err)
			break
		}
		s.handle(msg)
	}

	return s.Shutdown()
}

// Shutdown the server and clean up the socket
func (s *Server) Shutdown() error {
	info("shutting down")
	return s.Close()
}

//===========================================================================
// Message Handling
//===========================================================================

func (s *Server) handle(message *pb.Message) {
	info("received: %s\n", message.String())

	switch message.Type {
	case pb.MessageType_SINGLE:
		// Broadcast the single message to all peers
		s.net.Broadcast(message.Message)

		// Send a reply to the client
		reply := fmt.Sprintf("reply msg #%d", s.nRecv)
		s.send(reply, pb.MessageType_SINGLE)
	case pb.MessageType_BOUNCE:
		s.send("ACK", pb.MessageType_BOUNCE)
	}

}
