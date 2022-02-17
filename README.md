## Finnomena Recruitment Test
---

As a user running the application

I can view a list of funds in a user submitted time range (e.g. 1D, 1W, 1M, 1Y)

So that I know which funds are ranked by performance in the selected time range
<p>&nbsp;</p>  
 
### How to run:
&nbsp;  

Go to finnomena folder then run command :
```
go run main.go -t 1Y
```

or Build it and run command :
```
go build
./finnomena -t 1Y
```

The result will show on your command line.

&nbsp; 
### Usage of app:
```
-t string
    time range (e.g. 1D, 1W, 1M, 1Y) you want to view a list of funds (default "1Y")
```



