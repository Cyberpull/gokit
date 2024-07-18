package cyb

func sendErrorChan(err error, errChan ...chan error) {
	for _, e := range errChan {
		e <- err
	}
}
