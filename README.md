<h1 text-align="center"> Manaphy <img src="https://user-images.githubusercontent.com/49465310/124341427-3e168380-dbf7-11eb-960b-77cbd22b499c.gif" width="50"> </h1>

このレポジトリ（Manaphy）はバックエンドはGolang　フロントエンドはNUXT（VUE）で書かれています
This repository is written by golang for backend and Nuxt for frontend.

業務で行ったものをアウトプットや練習として書いてあるものになります
this code is going to write what I learned from my job.

そちらをご了承の上、お読みください！
you can read this code on this understanding:)


Backend (バックエンド)
そしてGolangも最近（2020/09~）から使い始めました！
Docker化をしている

【構成】
クリーンアーキテクチャを使っています
internal以下にdomain,handler,repository,usecase,interface等があります。各々、クリーンアーキテクチャでやることをやってますwほぼ本をみながらですがw

Frontend (フロントエンド)
最近（2021/05~）での業務でNuxtを使い、フロントエンドを書きました！アウトプットをするためにNuxtで書きますw

ここではスクレイピングで取ってきたおすすめ等の映画をデータベースから持ってきて表示する予定ですが（未定w）今のところモックでIndex.vueで表示をしている感じになります。

言わずもがな、Create-nuxt-app 等をすればこのディレクトリになります。そして主要なページはpages以下に、componentディレクトリには使いまわせるように中身を分けています！
