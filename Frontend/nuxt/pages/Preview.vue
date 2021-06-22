<template>
  <div>
    <ul class="text-center">
      <li v-for="post in posts">
      {{ post.id }} - {{ post.title }}
       <link-prevue v-bind:url= "post.title">
        <template slot-scope="props">
            <div>
              <div class="text-center">
                <img class="w-1/5" :src="props.img" :alt="props.title" />
              </div>
              <div class="card-block">
                <p class="text-lg text-center">{{props.title}}</p>
                <div class="text-center">
                  <a v-bind:href="props.url">More</a>
                </div>
              </div>
            </div>
          </template>
        </link-prevue>
      </li>
    </ul>
  </div>
</template>

<script>
import LinkPreview from 'link-prevue'

export default {
  components: {
    LinkPreview
  },
  asyncData(){
    const posts = require('~/assets/Postmock.json')
    return{
      posts
    }
  }
}
</script>
