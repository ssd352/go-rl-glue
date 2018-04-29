package main

import(
	"math/rand"
	// "testing"
	// "fmt"
)

// type Agent struct{
// 	value_function []float64
// }



// var(
// 	last_action []int
// 	num_actions = 10
// 	Value_function = []float64{0.1, 0.2, 0.3}
// )

type SimpleAgent struct{
	last_action []int
	num_actions int  // 10
	Value_function []float64  // {0.1, 0.2, 0.3}
}

func (agent *SimpleAgent) Init(){
	agent.num_actions = 10
	agent.Value_function = []float64{0.1, 0.2, 0.3}
	agent.last_action = []int{0}
}

func (agent *SimpleAgent) Start(this_observation []int) []int{
	agent.last_action[0] = rand.Intn(agent.num_actions)
	local_action := []int{0}
	local_action[0] = rand.Intn(agent.num_actions)
	return local_action
}

func (agent *SimpleAgent) Step(reward float64, this_observation []int) []int{
	local_action := []int{0}
	local_action[0] = rand.Intn(agent.num_actions)
	agent.last_action = local_action
	return agent.last_action
}

func (agent SimpleAgent) End(reward float64){
}

func (agent SimpleAgent) Cleanup(){
}

// TODO
func (agent SimpleAgent) Message(inMessage string){

}

// func main(){
// 	Init()
// 	fmt.Println(last_action)

// 	tmp := Start([]int{1})
// 	fmt.Println(last_action)
// 	fmt.Println(tmp)

// 	for i := 0; i < 10; i++ {
// 		tmp := Step(2, []int{1})
// 		fmt.Println(last_action)
// 		fmt.Println(tmp)
// 	}

// 	End(3)
// 	fmt.Println(last_action)
// }