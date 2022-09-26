package erratum

func Use(opener ResourceOpener, input string) (err error) {
	resource, err := opener()
	defer func() {
		if resource != nil {
			resource.Close()
		}
	}()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = opener()
	}

	defer func()  {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
				err = frobErr
			} else {
				err = r.(error)
			}
		}
	}()
	resource.Frob(input)
	return 
}
