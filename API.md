**Account API**
----
  Create user

* **URL**

  /user

* **Method:**
  
  `POST`

*  **Headers**

   **Required:**
 
   `X-Token-Gen=[string]`

* **Body Params**

    **Required:**

    `{
        email:[string],
        password:[string]
    }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ uuid : [string] }`
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Log in" }`

  OR

  * **Code:** 422 UNPROCESSABLE ENTRY <br />
    **Content:** `{ error : "Email Invalid" }`

----
  Get user

* **URL**

  /user/{uuid}

* **Method:**
  
  `GET`

*  **Headers**

   **Required:**
 
   `X-Token-Gen=[string]`

* **URL Params**

    **Required:**

    `{
        uuid:[string]
    }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ email : [string] }`
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Log in" }`

  OR

  * **Code:** 422 UNPROCESSABLE ENTRY <br />
    **Content:** `{ error : "Some error" }`
