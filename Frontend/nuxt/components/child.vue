<template>
  <div class="child">
    <link-prevue v-bind:url= "post.url">
        <template slot-scope="props">
            <div class="text-xs">
              <div class="relative flex justify-center ">
                <img class="w-1/4" :src="props.img" :alt="props.title" />
              </div>
              <div class="card-block">
                <div>
                  <p class="text-white text-xs text-center">{{props.title}}</p>
                </div>
                <div class="text-center text-white">
                  <a v-bind:href="props.url">More</a>
                </div>
              </div>
            </div>
          </template>
        </link-prevue>
  </div>
</template>

<script>
  export default {
    props: {
      post: {
        type: Object,
        required: true
      }
    }
  }
</script>
