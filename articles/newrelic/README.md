New Relic Example in GoLang with Deep Instrumentation
It is easy to integrate newrelic in GoLang. We just need one of the middleware to use the newrelic.Application. For eg in GIN framework of GO we can do something like this 
nrConfig := newrelic.NewConfig("test", "somekey")
nrapp, err = newrelic.NewApplication(nrConfig)
r := gin.Default()
r.Use(nrgin.Middleware(nrapp))
With this change you will able to see your application in new relic something like this 

But you will not be able to see breakdown of  api in "Transaction" because  Go is a compiled language . Hence unlike JAVA to see the breakup of the api you have to do explicit deep instrumentation in golang. Below is a simple example of how to use newrelic in the GIN web framework of go with deep instrumentation and using context.