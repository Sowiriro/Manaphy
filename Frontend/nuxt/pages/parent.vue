<template>
  <div>
    <ul class="text-center">
      <li v-for="post in posts">
      {{ post.id }} - {{ post.title }}
      <child :post="post" />
      </li>
    </ul>
  </div>
</template>

<script>
import child from '~/components/child.vue'

export default {
  components: {
    child
  },
  asyncData(){
    const posts = require('~/assets/Postmock.json')
    return{
      posts
    }
  }
}
</script>

