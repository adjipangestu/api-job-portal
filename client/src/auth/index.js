const API_URL = 'http://local.detik.com:8000'
const LOGIN_URL = API_URL + '/login'

export default {
    login(context, creds, redirect) {
        context.$http.post(LOGIN_URL, creds).then(response => {
            localStorage.setItem('token', response.body.token)
            console.log(response)
            if (redirect) {
                context.$router.replace(redirect)
            }
        }, response => {
            context.error = response.body.error
            console.log(response)
        })
    },

    logout(context) {
        localStorage.removeItem('token')
        context.$router.replace('/home')
    },

    isAuthenticated() {
        var jwt = localStorage.getItem('token')
        if (jwt) {
            return true
        }
        return false
    },

    getAuthHeader() {
        return {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        }
    }

}