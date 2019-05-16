# login

## 登陆

#### 接收

```json
post
json
{
    "id",       // string 学号
    "password"  // string 密码
}
```

#### 返回

```json
json
{
    "message",   // string "ok" / "user not exists" / "wrong password" / "fail"
    "identity",  // string "student" / "admin" / "teacher",
    "token":{   // string
        "id",
        "identity"
    }
}
```

# Student

## Profile

### 默认

#### 接收

```json
Get
Header
{
    "Authorization" // token
}
```

#### 返回

```json
json
{
    "message",  // string "ok" / "fail" / "invalid token"
    "id",       // string 学号
    "gender",   // string 性别
    "name",     // string 姓名
    "dname",    // string 院系名
}
```

## Course Calendar

### 查询已选课程（学期)

#### 接收

```json
Post
Header
{
    "Authorization"
}
json
{
    "term",  // 学期
    "op": "query"
}
```

#### 返回

```json
json
{
    "message",    // string "ok"  / "fail"
    "cid",       // []string 课程号
    "cname",     // []string 课程名
    "tid",       // []string 教师号
    "tname",     // []string 教师名
    "credit",    // []string 学分
    "classTime"  // []string 上课时间
}
```

### 退课

#### 接收

```json
Post
Header
{
    "Authorization"
}
json
{
    "cid",    // 课程号
    "term",   // 学期
    "op": "delete"
}
```

#### 返回

```json
json
{
    "message"    // string "ok" / "cid not exits" / "fail"
    "cid",       // []string 课程号
    "cname",     // []string 课程名
    "tid",       // []string 教师号
    "tname",     // []string 教师名
    "credit",    // []string 学分
    "classTime"  // []string 上课时间
}
```

## Couser Query

### 查询课程

#### 接收

```json
Post
json
{
    // 全空就是查所有课程
    "cid": "",  // 默认为空值
    "tid": "",
    "cname": "",
    "tname": "",
    "op" : "query",
    "Authorization"
}
```

#### 返回

```json
json
{
    "message",  // string "ok" / "fail" 
    "cid",      // []string
    "cname",    // []string
    "tid",      // []string
    "tname",    // []string
    "credit",   // []string
    "classTime" // []string
}
```

### 选课

#### 接收

```json
Post
json
{
    "term",
    "classTime",
    "cid",
    "tid",
    "op": "select",
    "Authorization"
}
```

#### 返回

```json
json
{
    "message"   // string "ok" / "fail"
}
```

## Score Table

### 选择学期

#### 接收

```json
Post
Header
{
    "Authorization"
}
Body
{
    "term"
}
```

#### 返回

```json
json
{
    "message",      // string "ok" / "fail"
    "cid",          // []string
    "cname",        // []string
    "tid",          // []string
    "tname",        // []string
    "credit",       // []string
    "score"         // []string
}
```

## Score Summary

### 默认

#### 接收

```json
Get
Header
{
    "Authorization"
}
json
{
    "term" 
}
```

#### 返回

```json
{
    "message",      // string "ok" / "fail"
    "cid",          // []string
    "cname",        // []string
    "credit",       // []string
    "score",        // []string
    "term"          // []string
}
```

## Score Trend

### 默认

#### 接收

```json
Get
Header
{
    "Authorization"
}
json
{
    "term" // []string
}
```

#### 返回

```json
{
    "score" 	// []string
}
```



## Score College

### 默认

#### 接收

```json
Get
Header
{
    "Authorization"
}
```

#### 返回

```json
{
    "id",		// sting
    "name",		// string
    "score",	// string
    "dname",	// string
    "totalStudents",	// string
    "rank",				// string
    "percentage"		// string
}
```



# Teacher

## Profile

### 默认

#### 接收

```json
Get
Header
{
    "Authorization" // token
}
```

#### 返回

```json
json
{
    "message",  // string "ok" / "fail" / "invalid token"
    "tid",       // string 工号
    "gender",   // string 性别
    "tname",     // string 姓名
    "dname",    // string 院系名
}
```

## Score Management

### 默认

#### 接收

```json
GET
Header
{
    "Authorization"
}
```

#### 返回

```json
json
{	// 教师所教的课程
    "cid",		// []string
    "cname",	// []string
    "term"		// []string
}
```



### 选择课程

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
    "cid",
    "term",
    "op" = "query"
}
```

#### 返回

```json
json
{
    "message",	// string "ok" / "fail
    "id",		// []string
    "name",		// []string
    "dname",	// []string
    "score"		// []string
}
```

### 修改成绩

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
    "id",
    "cid",
    "score",	// string 所修改的分数
    "op" = "change"
}
```

#### 返回

```json
json
{
    "message" // string "ok" / "fail"
}
```

## Score Analysis

### 选择课程

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
    "cid",
    "term"
}
```

#### 返回

```json
json
{
    "message",	// string "ok" / "fail
    "score"		// []string
}
```



## Course Calendar

### 选择学期

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	
    "term"
}
```

#### 返回

```json
json
{	// 该教师每学期所教的课
    "message",	// string "ok"/"fail"
    "cid",		// []string
    "cname",	// []string
    "credit",	// []string
    "classTime"	// []string
}
```

## Class Table

### 选择课程 

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	
    "term",
    "cid"
}
```

#### 返回

```json
json
{
    "message",
    "id",			// []string	
    "name",			// []string
    "gender",		// []string
    "birthplace",	// []string
    "phone",		// []string
    "dname"			// []string
}
```

# admin

## Profile

### 默认

#### 接收

```json
GET
Header
{
    "Authorization"
}
```

#### 返回

```json
{
    "aid"
}
```

## Teacher Management

### 新增

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
    "tid",
    "tname",
    "gender",
    "birthday",
    "wage",
    "education",
    "dname",
    "op": "add"
}
```

#### 返回

```json
json
{
    "message"	// "ok" / ”fail"
}
```

### 修改

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	// 默认为空，表示不修改这一项
    "tid", 	// tid 不为空，且不能修改
    "tname",
    "gender",
    "birthday",
    "wage",
    "education",
    "dname",
    "op": "modify"
}
```

#### 返回

```json
json
{
    "message"	// "ok" / ”fail"
}
```



### 删除

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	
    "tid",
   	"op": "delete"
}
```

#### 返回

```json
json
{
    "message"	// "ok" / ”fail"
}
```

## Student Management

### 新增

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
    "id",
    "name",
    "gender",
    "birthday",
    "birthplace",
    "phone",
    "dname",
    "op": "add"
}
```

#### 返回

```json
json
{
    "message"	// "ok" / ”fail"
}
```

### 修改

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	// 可以不传，也可以为空
    "id", 	// id 不为空，且不能修改
    "name",
    "gender",
    "birthday",
    "birthplace",
    "phone",
    "dname",
    "op": "modify"
}
```

#### 返回

```json
json
{
    "message"	// "ok" / ”fail"
}
```



### 删除

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	
    "id",
   	"op": "delete"
}
```

#### 返回

```json
json
{
    "message"	// "ok" / ”fail"
}
```

## Course Management(!)

### 默认

#### 接收

```json
GET
Header
{
    "Authorization"
}
```

#### 返回

```json
json
{
    "message",
    "term"	// []string
}
```

### 选择学期

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	"message",
    "term",		// string
 	"op":"select"
}
```

#### 返回

```json
json
{
    "message",	// string
    "cid",		// []string
    "cname",	// []string
    "credit", 	// []string
    "dname",	// []string
    "tid",		// []string
    "tname",	// []string
    "classTime"	// []string
}
```

### 新增

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{	
 	"cid",
 	"cname",
 	"tid",
 	"credit",
 	"classTime",
 	"did",
    "term",
    "op":"add"
}
```

#### 返回

```json
json
{
    "message"	// string "ok" / "fail"
}
```



### 修改

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
    // 可以不传，也可以为空值
    "message",
 	"cid",			// 不为空，且不能修改
 	"cname",
 	"tid",
 	"credit",
 	"classTime",
 	"did",
    "term",
    "op":"modify"
}
```

#### 返回

```json
json
{
    "message"	// string "ok" / "fail"
}
```



### 删除

#### 接收

```json
POST
Header
{
    "Authorization"
}
json
{
 	"cid",			
    "tid",
    "term",
    "op":"delete"
}
```



#### 返回

```json
json
{
    "message"	// string "ok" / "fail"
}
```

