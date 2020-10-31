package main

import  (
    "fmt"
    "math/rand"
)

type Job struct {
    Id int
    Random int
}


type Result struct {
    Job *Job
    Sum int
}

// 通过多个线程来完成多个任务
func main(){
  // 设置任务通道
  jobChan :=make(chan *Job,40)
  // 设置结果通道
  resultChan := make(chan *Result ,40)
  createWorker(60,jobChan,resultChan)

  go func (resultChan <-chan *Result){
        for result := range resultChan {
            fmt.Printf("job :%d,random:%d,sum:%d\n",result.Job.Id,result.Job.Random,result.Sum)
        }
  }(resultChan)


  for i:=1;;i++{
        job := &Job{
            Id : i,
            Random : rand.Intn(100),
        }
       jobChan <- job
  }
}

func createWorker(num int ,jobChan  <-chan *Job, resultChan  chan<- *Result) {
    for i :=0 ; i <num ; i++{
        go func (jobChan <-chan *Job, resultChan  chan<- *Result){
             for job := range jobChan {
                sum := 0
                tmp := job.Random
                 for tmp>0 {
                    sum +=tmp%10
                    tmp = tmp/10
                 }
                 result := &Result{
                   Job : job,
                   Sum : sum,
                 }
                 resultChan <-  result
             }
        }(jobChan, resultChan)
    }
}



