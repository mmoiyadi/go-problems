package main

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
type reactor struct {
}


type canceler struct{
	cell *cell
	callbackIndex int
}

type callback_struct struct {
	callback func(int) 
	isValid bool
}


type cell struct{
	val int

	depCell1 *cell
	depCell2 *cell

	/// Newer Implementation
	changed bool

	depCallbacks []func(int)

	callbacks []callback_struct
}

func (c *cell) IsCompute1() bool{
	return c.depCell1 != nil && c.depCell2 == nil
}

func (c *cell) IsCompute2() bool{
	return c.depCell1 != nil && c.depCell2 != nil
}

func (c *cell) SetCallback( callback func(int) ){
	if c != nil{
		if c.IsCompute1(){
			cellCb := c.depCell1
			cellCb.SetCallback( callback)
		}else if c.IsCompute2(){
			cellCb1 := c.depCell1
			cellCb2 := c.depCell2
			cellCb1.SetCallback( callback)
			cellCb2.SetCallback( callback)
		}else {
				c.depCallbacks = append(c.depCallbacks, callback)
		}
	}
	
}

func (c *canceler) Cancel() {
	 if c.cell != nil{
		for idx := range c.cell.callbacks{
			if idx == c.callbackIndex{
				c.cell.callbacks[idx].isValid = false
			}
		}
	 }
}

func (c *cell) Value() int {
	return c.val
}

func (c *cell) SetValue(value int) {
		if value == c.val{
			c.changed = false
		} else{
			c.changed = true
		}
		c.val = value

		computedVal := value

		if len(c.depCallbacks) != 0{
			for _, cb := range c.depCallbacks{
				cb(computedVal)
			}
		}
		
		if len(c.callbacks) != 0{
			for _, cb_struct := range c.callbacks{
				if cb_struct.isValid{
					cb_struct.callback(computedVal)
				}
			}
		}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	canc := &canceler{}
	if c.depCell1 != nil{
		var f = func(n int){
			if c.changed{
				callback(c.Value())			
			}
		}
		
		depC := c.depCell1
		for depC.depCell1 != nil{
			depC = depC.depCell1
		}

		depC.callbacks = append(depC.callbacks, callback_struct{callback: f, isValid: true})

		canc.callbackIndex = len(depC.callbacks) - 1
		canc.cell = depC
	}
	return canc
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	inputCell := cell{val: initial}
	return &inputCell
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	computeCell := cell{ changed: false}
	c, ok := dep.(*cell)
	if !ok {
		panic("Sorry about that!!!")
	}
	
	var f = func(n int){
		computeCell.SetValue(compute(c.Value()))
	
	}
	c.SetCallback(f)

	computeCell.depCell1 = c
	computeCell.SetValue(compute(c.Value()))
	//computeCell.changed = false
	return &computeCell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	computeCell := cell{changed: false}
	c1, ok := dep1.(*cell)
	if !ok {
		panic("Sorry about that!!!")
	}

	c2, ok := dep2.(*cell)
	if !ok {
		panic("Sorry about that!!!")
	}
	
	var f = func(n int){
		computedVal := compute(c1.Value(),c2.Value())
		if computedVal != computeCell.val{
			computeCell.SetValue(computedVal)
		}
		
	}
	
	c1.SetCallback(f)
	c2.SetCallback(f)
	
	computeCell.depCell1 = c1
	computeCell.depCell2 = c2
	computedVal := compute(c1.Value(),c2.Value())
	computeCell.SetValue(computedVal)
	computeCell.changed = false
	return &computeCell
}
