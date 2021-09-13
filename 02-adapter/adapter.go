package adapter

// Motor 定义目标接口
type Motor interface {
	Drive() string
}


// ElectricMotor 适配者1：电能发动机
type electricMotor struct {}

func (e *electricMotor) electricDrive() string {
	return "electricDrive"
}

func NewElectricMotor() *electricMotor {
	return &electricMotor{}
}

// NewElectricAdapter 是 electricAdapter的工厂函数 将 被适配对象转换为 目标接口
func NewElectricAdapter(adapted *electricMotor) Motor {
	return &electricAdapter{
		electricMotor: *adapted,
	}
}

// electricAdapter 将 被适配对象转换为 目标接口的适配器
type electricAdapter struct {
	electricMotor
}

// Drive 具体实现适配的方法
func (e *electricAdapter) Drive() string {
	return e.electricDrive()
}

// opticalMotor 适配者2：光能发动机
type opticalMotor struct {}

func (e *opticalMotor) opticalDrive() string {
	return "opticalDrive"
}

func NewOpticalMotor() *opticalMotor {
	return &opticalMotor{}
}

// NewOpticalAdapter 是 opticalMotor 的工厂函数 将 被适配对象转换为 目标接口
func NewOpticalAdapter(adapted *opticalMotor) Motor {
	return &opticalAdapter{
		opticalMotor: *adapted,
	}
}

// opticalAdapter 将 被适配对象转换为 目标接口的适配器
type opticalAdapter struct {
	opticalMotor
}

// Drive 具体实现适配的方法
func (e *opticalAdapter) Drive() string {
	return e.opticalDrive()
}

func MotorAdapter(name string) Motor {
	if name == "electric" {
		electricMotor := NewElectricMotor()
		adapter := NewElectricAdapter(electricMotor)
		return adapter

	} else if name == "optical" {
		opticalMotor := NewOpticalMotor()
		adapter := NewOpticalAdapter(opticalMotor)
		return adapter
	}
	return nil
}

