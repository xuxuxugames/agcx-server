# API

## Error

    {
        "status": 4xx/5xx,
        "message": "error message"
    }

## User

### 用户认证

POST `/user/auth`

#### JSON 参数

 - email: 登陆账号
 - password: 用户密码

#### 响应
    {
        "status": 200,
        "token": "{Token}"
    }

### 注册用户

POST `/user`

#### JSON 参数

 - name: 昵称
 - email: 登陆邮箱
 - password: 密码

#### 响应

    {
        "status": 201,
        "user_id": 2
    }

### 修改密码

PUT `/user/{user_id}/password`

#### Header

 - Authorization: Bearer {Token}

#### JSON 参数

 - old_password: 原密码
 - new_password: 新密码

#### 响应

    {
        "status": 200
    }

## Score

### 保存成绩

POST `/score`

#### Header

 - Authorization: Bearer {Token}

#### JSON 参数

 - user_id 用户id
 - game 游戏名称
 - score 本次成绩

#### 响应

    成功：
    {
        "status": 201,
    }
    
### 获取游戏成绩列表

GET '/score'

#### URL Query 参数

 - user_id: 用户id
 - game: 游戏名称
 - page: 分页
 - start_at: 开始时间
 - end_at: 结束时间
 
#### 响应
 
    {
        "status": 200,
        "total": 100,
        "current_page": 1,
        "per_page": 20,
        "data": [
            {
                "id": 1,
                "user_id": 1,
                "game": '2048',
                "score": 1,
                "ip": 
            }
            ...
        ]
    }
