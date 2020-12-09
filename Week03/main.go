package main

func main(){
  //初始化路由
	router := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello")
	});

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 9000),
		Handler:        router,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: 1 << 20,
	}
	group, c := errgroup.WithContext(context.Background())
	//启动http服务
	group.Go(func() error {
		fmt.Printf("Http server run —— Actual pid is %d\n", syscall.Getpid())
		if e :=  server.ListenAndServe(); e != nil{
			fmt.Println(e)
			return e
		}
		return nil
	})

	// 注册监听linux signal信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// 其他任务运行
	for i := 0; i < 10; i++ {
		index := i
		time.Sleep(time.Duration(index) / 2  * time.Second)
		group.Go(func() error {
			select {
			case <-sigs:		// 监听signal关闭信号
				fmt.Printf("task%d：注销（signal信号） \n", index)
				sc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				server.Shutdown(sc)
			case <- c.Done():	// 如果其他任务异常中断，则当前任务注销
				fmt.Printf("task%d：注销 \n", index)
				return c.Err()
			default:
			}
			fmt.Printf("task %d run success \n", index)
			return nil
		})
	}

	if e := group.Wait(); e != nil {
		fmt.Println(fmt.Sprintf("first error is %v", e))
	} else {
		fmt.Println("http server close")
	}

}
