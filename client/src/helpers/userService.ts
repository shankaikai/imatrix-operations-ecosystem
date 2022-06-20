import bcrypt from 'bcryptjs';
import crypto from 'crypto'
export const isLoggedIn = () => {
    //TODO: do some validation for expired jwt
    return localStorage.getItem('jwt') === "JWT"
}

export const storeJWT = (jwt:string) => {
    localStorage.setItem('jwt',jwt)
}
export const signIn = (username:string, password:string) : boolean => {
    //TODO: retrieve random string from BE
    const randomString = "abcde"
    //our key is the password.
    const hash = crypto.createHmac('sha256',password).update(randomString).digest('hex')

    //TOOD: send back to backend for validation
    const jwt = 'JWT'

    if (jwt) {
        storeJWT(jwt)
        return true
    } else {
        return false
    } 
}