# 博客网站 RESTful API 设计

## 总览

以下 `url` 默认将 `https://api.myblog.com/` 简写为 `/api/` ，如 `https://api.myblog.com/users` 简写为 `/api/users` 。

| URL                                               | HTTP   | 功能           | 鉴权 |
| ------------------------------------------------- | ------ | -------------- | ---- |
| `/api/captcha`                                    | GET    | 获取图形验证码 | ×    |
| `/api/users`                                      | POST   | 用户注册       | ×    |
| `/api/login`                                      | POST   | 用户登录       | ×    |
| `/api/users/{username}`                           | GET    | 查看用户       | ×    |
| `/api/users/{username}/articles`                  | GET    | 获取文章列表   | ×    |
| `/api/users/{username}/articles/{title}`          | GET    | 获取文章       | ×    |
| `/api/users/{username}/articles/{title}/comments` | GET    | 获取评论列表   | ×    |
| `/api/users/{username}/articles/{title}/comments` | POST   | 添加评论       | √    |
| `/api/users/{username}/articles`                  | POST   | 发布文章       | √    |
| `/api/users/{username}/articles`                  | PUT    | 修改文章       | √    |
| `/api/users/{username}/articles`                  | DELETE | 删除文章       | √    |
| `/api/users/password`                             | PUT    | 修改密码       | √    |
| `/api/users/logout`                               | POST   | 登出           | √    |

## 获取图形验证码

- API： `GET /api/captcha`

- Request

    ```json
    {
        "version": "1"
    }
    ```


- Response

    ```json
    {
        "captchaId": "18",
    	"captcha": "a8we15af8e43qaw2e1r54wer5445we..."
    }
    ```

## 用户注册

提供用户名、密码、验证码 id 以及用户输入的验证码结果。

- API： `POST /api/users`

- Request

    ```json
    {
        "version": "1",
        "username": "liangwj",
        "password": "123456",
        "captchaId": "23",
        "captchaCode": "a8we15af8e43qaw2e1r54wer5445we..."
    }
    ```


- Response

    ```json
    {
        "username": "liangwj",
        "id": 12345678
    }
    ```


## 用户登录

- API： `POST /api/login`

- Request

    ```json
    {
        "version": "1",
        "username": "liangwj",
        "password": "123456",
        "captchaId": "123",
        "captchaCode": "a8we15af8e43qaw2e1r54wer5445we..."
    }
    ```


- Response

    ```json
    {
        "username": "liangwj",
        "id": 12345678
    }
    ```

## 查看用户

- API： `GET /api/users/{username}`

- Request

    ```json
    {
        "version": "1"
    }
    ```


- Response

    ```json
    {
        "login": "{current_user}",
        "id": 12345678,
        "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
        "html_url": "https://myblog.com/{username}",
        "articles_url": "/api/users/{username}/articles",
        "subscriptions_url": "/api/users/{username}/subscriptions",
        "fans_url": "/api/users/{username}/fans",
        "subscriptions": 58,
        "fans": 12,
        "created_at": "2017-09-22T15:07:28Z",
        "updated_at": "2019-11-18T19:56:21Z"
    }
    ```

## 获取文章列表

- API： `GET /api/users/{username}/articles`

- Request

    ```json
    {
        "version": "1"
    }
    ```


- Response

    ```json
    [
        {
            "id": 155829275,
            "title": "{title}",
            "owner": {
                "login": "{current_user}",
                "id": 12345678,
                "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
                "html_url": "https://myblog.com/{username}",
                "articles_url": "/api/users/{username}/articles",
                "subscriptions_url": "/api/users/{username}/subscriptions",
                "fans_url": "/api/users/{username}/fans",
                "subscriptions": 58,
                "fans": 12,
                "created_at": "2017-09-22T15:07:28Z",
                "updated_at": "2019-11-18T19:56:21Z"
            },
            "html_url": "https://myblog.com/{username}/{title}",
            "description": "{description}",
            "url": "/api/users/{username}/articles/{title}",
            "comments_url": "/api/users/{username}/articles/{title}/comments",
            "tags_url": "/api/users/{username}/articles/{title}/tags",
       	 	"content_url": "/api/users/{username}/articles/{title}/content",
            "like": 88,
        	"comment": 15,
            "created_at": "2018-11-02T07:23:36Z",
            "updated_at": "2019-02-07T15:31:05Z",
        }
    ]
    ```

## 获取文章

- API： `GET /api/users/{username}/articles/{title}`

- Request

    ```json
    {
        "version": "1"
    }
    ```


- Response

    ```json
    {
        "id": 155829275,
        "title": "{title}",
        "owner": {
            "login": "{current_user}",
            "id": 12345678,
            "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
            "html_url": "https://myblog.com/{username}",
            "articles_url": "/api/users/{username}/articles",
            "subscriptions_url": "/api/users/{username}/subscriptions",
            "fans_url": "/api/users/{username}/fans",
            "subscriptions": 58,
            "fans": 12,
            "created_at": "2017-09-22T15:07:28Z",
            "updated_at": "2019-11-18T19:56:21Z"
        },
        "html_url": "https://myblog.com/{username}/{title}",
        "description": "{description}",
        "url": "/api/users/{username}/articles/{title}",
        "comments_url": "/api/users/{username}/articles/{title}/comments",
        "content_url": "/api/users/{username}/articles/{title}/content",
        "like": 88,
        "comment": 15,
        "created_at": "2018-11-02T07:23:36Z",
        "updated_at": "2019-02-07T15:31:05Z",
    }
    ```

## 获取评论列表

- API： `GET /api/users/{username}/articles/{title}/comments`

- Request

    ```json
    {
        "version": "1"
    }
    ```


- Response

    ```json
    [
        {
            "url": "/api/users/{username}/articles/{title}/comments/15",
            "article_url": "/api/users/{username}/articles/{article_title}",
            "id": 526654256,
            "number": 15,
            "title": "{title}",
            "user": {
                "login": "{current_user}",
                "id": 12345678,
                "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
                "html_url": "https://myblog.com/{username}",
                "articles_url": "/api/users/{username}/articles",
                "subscriptions_url": "/api/users/{username}/subscriptions",
                "fans_url": "/api/users/{username}/fans",
                "subscriptions": 58,
                "fans": 12,
                "created_at": "2017-09-22T15:07:28Z",
                "updated_at": "2019-11-18T19:56:21Z"
            },
            "created_at": "2019-11-21T15:08:26Z",
            "updated_at": "2019-11-21T16:12:10Z",
            "body": "xxxx"
        }
    ]
    ```

## 添加评论

需要鉴权。

- API： `POST /api/users/{username}/articles/{title}/comments`

- Request

    ```json
    {
        "version": "1",
        "article_id": 4984654,
        "user_id": 12345678,
        "body": "xxx"
    }
    ```


- Response

    ```json
    {
        "id": 155829275,
        "title": "{title}",
        "owner": {
            "login": "{current_user}",
            "id": 12345678,
            "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
            "html_url": "https://myblog.com/{username}",
            "articles_url": "/api/users/{username}/articles",
            "subscriptions_url": "/api/users/{username}/subscriptions",
            "fans_url": "/api/users/{username}/fans",
            "subscriptions": 58,
            "fans": 12,
            "created_at": "2017-09-22T15:07:28Z",
            "updated_at": "2019-11-18T19:56:21Z"
        },
        "html_url": "https://myblog.com/{username}/{title}",
        "description": "{description}",
        "url": "/api/users/{username}/articles/{title}",
        "comments_url": "/api/users/{username}/articles/{title}/comments",
        "content_url": "/api/users/{username}/articles/{title}/content",
        "like": 0,
        "comment": 0,
        "created_at": "2018-11-02T07:23:36Z",
        "updated_at": "2019-02-07T15:31:05Z",
    }
    ```

## 发布文章

需要鉴权。

- API： `POST /api/users/{username}/articles`

- Request

    ```json
    {
        "version": "1",
        "title": "{title}",
        "description": "xxx",
        "user_id": 12345678,
        "content": "xxx"
    }
    ```


- Response

    ```json
    {
        "id": 155829275,
        "title": "{title}",
        "owner": {
            "login": "{current_user}",
            "id": 12345678,
            "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
            "html_url": "https://myblog.com/{username}",
            "articles_url": "/api/users/{username}/articles",
            "subscriptions_url": "/api/users/{username}/subscriptions",
            "fans_url": "/api/users/{username}/fans",
            "subscriptions": 58,
            "fans": 12,
            "created_at": "2017-09-22T15:07:28Z",
            "updated_at": "2019-11-18T19:56:21Z"
        },
        "html_url": "https://myblog.com/{username}/{title}",
        "description": "{description}",
        "url": "/api/users/{username}/articles/{title}",
        "comments_url": "/api/users/{username}/articles/{title}/comments",
        "content_url": "/api/users/{username}/articles/{title}/content",
        "like": 0,
        "comment": 0,
        "created_at": "2018-11-02T07:23:36Z",
        "updated_at": "2019-02-07T15:31:05Z",
    }
    ```

## 修改文章

需要鉴权。

- API： `PUT /api/users/{username}/articles`

- Request

    ```json
    {
        "version": "1",
        "title": "{title}",
        "description": "xxx",
        "userID": 12345678,
        "content": "xxx"
    }
    ```


- Response

    ```json
    {
        "id": 155829275,
        "title": "{title}",
        "owner": {
            "login": "{current_user}",
            "id": 12345678,
            "avatar_url": "https://avatars.myblogusercontent.com/u/12345678?v=4",
            "html_url": "https://myblog.com/{username}",
            "articles_url": "/api/users/{username}/articles",
            "subscriptions_url": "/api/users/{username}/subscriptions",
            "fans_url": "/api/users/{username}/fans",
            "subscriptions": 58,
            "fans": 12,
            "created_at": "2017-09-22T15:07:28Z",
            "updated_at": "2019-11-18T19:56:21Z"
        },
        "html_url": "https://myblog.com/{username}/{title}",
        "description": "{description}",
        "url": "/api/users/{username}/articles/{title}",
        "comments_url": "/api/users/{username}/articles/{title}/comments",
        "content_url": "/api/users/{username}/articles/{title}/content",
        "like": 0,
        "comment": 0,
        "created_at": "2018-11-02T07:23:36Z",
        "updated_at": "2019-02-07T15:31:05Z",
    }
    ```

## 删除文章

需要鉴权。

- API： `DELETE /api/users/{username}/articles`

- Request

    ```json
    {
        "version": "1",
        "id": 5448441,
        "userID": 12345678,
    }
    ```


- Response

    ```json
    {
        
    }
    ```

## 修改密码

需要鉴权。

- API： `PUT /api/users/password`

- Request

    ```json
    {
        "version": "1",
        "id": 12345678,
        "username": "liangwj",
        "password": "123456",
        "new_password": "12345678",
        "captchaCode": "a8we15af8e43qaw2e1r54wer5445we..."
    }
    ```


- Response

    ```json
    {
        "username": "liangwj",
        "id": 12345678
    }
    ```

## 登出

需要鉴权。

- API： `POST /api/users/logout`

- Request

    ```json
    {
        "version": "1",
        "id": 12345678,
    }
    ```


- Response

    ```json
    {
        
    }
    ```

