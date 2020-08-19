# wiki go
golangで簡単なwikiを作成する

# データベース
db:wikigo
table:
article→id,title,body,updated_at,created_at

model

controllers

# url
/index
・全ての記事一覧を表示、それらのリンクも表示

/view/{{id}}
・その記事の内容を表示
・editへのリンク

/edit/{{id}}
・その記事の修正

/save/{{id}}
・修正案を反映