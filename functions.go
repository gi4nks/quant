package quant

type action0 func()

type action1 func() (interface {}, error)


func apply1(a action1) interface{} {
	ret, err := a()
	
	if err != nil {
		//parrot.Error("Error...", err)
		return nil
	} 
	
	return ret
}
