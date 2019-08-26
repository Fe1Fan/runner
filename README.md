# Runner [![Build Status](https://travis-ci.org/feifan00x/runner.svg?branch=master)](https://travis-ci.org/feifan00x/runner)
Shell Factory

### run:
```
❯ runner
__________                                  
\______   \__ __  ____   ____   ___________ 
 |       _/  |  \/    \ /    \_/ __ \_  __ \
 |    |   \  |  /   |  \   |  \  ___/|  | \/
 |____|_  /____/|___|  /___|  /\___  >__|   
        \/           \/     \/     \/ v 1.0.0
path check: true
file check: true
+-------+------+-------------+---------+--------+-----+
| INDEX | NAME |   REMARK    | VERSION | STATUS | PID |
+-------+------+-------------+---------+--------+-----+
|   1   | test | test remark |  1.0.0  |  stop  | nil |
+-------+------+-------------+---------+--------+-----+
input s scan config or index number exec.
```
### conf
```json
{
	"configs":[
		{
			"name": "test",
			"remark": "test remark",
			"ver": "1.0.0",
			"cmd": "echo helloworld",
			"incl": "",
			"status":"stop",
			"pid":"nil"
		}
	]
}
```