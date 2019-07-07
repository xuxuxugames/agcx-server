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

 - name: 用户姓名
 - password: 用户密码

#### 响应
    {
        "status": 200,
        "token": "{Token}"
    }

### 注册用户

POST `/user`

#### JSON 参数

 - name: 姓名
 - password: 密码

#### 响应

    {
        "status": 201,
        "user_id": 2
    }
    
### 删除用户

DELETE `/user/{user_id}`

#### Header

 - Authorization: Bearer {Token}

#### 响应

    {
        "status": 200
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
 - game_id 游戏id
 - score 本次成绩

#### 响应

    成功：
    {
        "status": 201,
        "message": "best score saved"
    }
    
    失败：
    {
        
    }
    
### 获取游戏成绩列表

GET '/score'

#### URL Query 参数

 - user_id: 用户id
 - game_id: 游戏id
 - page: 分页
 
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
                "game_id": 1,
                "score": 1,
            }
            ...
        ]
    }
