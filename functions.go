package quant

type Action0 func()

type Action1 func() (interface {}, error)


func apply1(a Action1) interface{} {
	ret, err := a()
	
	if err != nil {
		//parrot.Error("Error...", err)
		return nil
	} 
	
	return ret
}
