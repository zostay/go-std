package deferred

import "errors"

// Error is intended to be used in cases where you have a deferred function that
// returns an error and you want to capture that error, but you may have an error
// that has already occurred and need to capture that too:
//
//		func ProcessFile(name string) (err error) {
//			var r io.Reader
//	     r, err = os.Open(name)
//			if err != nil {
//				return
//			}
//			defer func() { deferred.Error(&err, r.Close()) }()
//
//			// process file
//			return
//		}
//
// Now, both errors can be captured without a lot of gymnastics.
func Error(err *error, deferErr error) {
	if deferErr == nil {
		return
	}

	if *err == nil {
		*err = deferErr
		return
	}

	*err = errors.Join(*err, deferErr)
}
