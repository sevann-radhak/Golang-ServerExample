package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/home", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.ListenServer()
}
