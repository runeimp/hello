using build
class Build : BuildPod
{
	new make()
	{
		podName = "hello_str"
		summary = "hello world pod"
		depends = ["sys 1.0"]
		srcDirs = [`fan/`]
	}
}