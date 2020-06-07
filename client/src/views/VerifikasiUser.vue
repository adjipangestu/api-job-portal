<template>
    <div>
        <p>{{ message }}</p>
    </div>
</template>
<script>
    import axios from "axios";
    import env from "../env.json";
    export default {
        name: "VerifikasiUser",
        data: () => ({
            message: '',
        }),

        mounted() {
            this.verifikasi()
        },
        methods: {
            verifikasi(){
                let token = this.$route.params.token
                axios.post( env.base_url_api + '/verifikasi/' + token, {withCredentials: true })

                    .then(response => {
                        //let data = response.data
                        let statusCode = response.status
                        let email = response.data.email
                        if (statusCode == 200) {
                            this.message = "Verifikasi "+ email +" berhasil"
                        }
                        console.log(response.data)
                    }).catch( error => {
                        let statusCode = error.response.status
                        if (statusCode == 500) {
                            this.message = error.response.data.error
                        }
                        console.log(error.response.data.error)
                })
            }
        }
    }
</script>
