<template>
<!---    最初は背景色を変える ---->
<div class= "p-6 md:p-16 bg-LightCodeBlue">
  <div>
    <h2 class="font-serif p-10 text-center text-2xl text-white"> Here is Index page </h2>
    <div class="p-5 text-white">
      <div class="text-right">
        <nuxt-link to="/Second"> Go to the Second Page</nuxt-link>
        <nuxt-link to="/parent"> Go to the parent Page</nuxt-link>
      </div>
      <div class="text-left">
        <nuxt-link to="/Preview"> Go to the Preview Page</nuxt-link>
      </div>
    </div>
    <ul class="p-4 flex justify-between mb-2 mr-1 flex-wrap">
      <li class="bg-CodeBlue mb-16 h-48 md:max-w-xs md:w-10/12 w-full h-16 rounded-lg p-10" v-for="post in posts">
        <child :post="post" />
      </li>
    </ul>
  </div>
</div>
</template>

<script>
import child from '~/components/child.vue'
import axios from 'axios'
export default {
  components: {
    child
  },
  data () {
    return {
      posts: []
    }
  },
  asyncData ({params}) {
    return axios
    // コレは例のURL　This is the example of google spread sheet url
    .get('https://script.googleusercontent.com/macros/echo?user_content_key=mpvmNDNkZDi5Wisso3xhs_zGS5bvdbZXPBruJqtYt4GUQ7RC5474K9GaWXRpsfgLUcKVtVR7nRmEM7EJOmTA-jf3mAhb9G7Im5_BxDlH2jW0nuo2oDemN9CCS2h10ox_1xSncGQajx_ryfhECjZEnMu-m7xoBSCuh7iMgVWWat2rcW2zI-2IsKmpJBUfHjMBh5b4Uu_yOQuF6F8QwzKSJu1Vp2z-ulPZYT-N0Y0xvQSB_zLlk9y9r9z9Jw9Md8uu&lib=Mu57dS6AvYCS9h_YriHoI4wrvUygKYYfT')
    .then((res) => {
      return {posts: res.data}
    })
  }
}
</script>
