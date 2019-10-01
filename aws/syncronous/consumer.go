package aws_syncronous

func (c *consumer)Consume() {
	for {
		out, err := receiveSQSMessages(c.QueueURL)

		if err != nil {
			continue
		}
	}

	for _, m := range out.Messages {

		c.run(newMessage(m))

		if err := h(m); err != nil {
			//	  log error
			continue
		}

		c.delete(m)
	}
}
