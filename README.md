A link to github with project that has 1 endpoint 
POST /user and accepts JSON with

username  //required
email //required, email
password //required 6-120 characters

and returns user_id in response

With (fake)repository which does a dummy check if such user already exists

What we'd like to see:
 - Good structured code
 - Error handling
 - input data validation
 - as many production-ready features as you can think of (graceful shutdown/logs/etc)