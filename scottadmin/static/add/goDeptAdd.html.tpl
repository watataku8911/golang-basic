<!DOCTYPE html>
<html lang="ja">
<head>
    <title>部門情報追加｜ScottAdmin</title>
    <link rel="stylesheet" type="text/css" href=" ../css/main.css">
    <meta charset="utf-8">
</head>
<body>
<h1>部門情報追加</h1>
<nav>
    <ul>
        <li><a href=" /">TOP</a></li>
        <li><a href=" /view">部門リスト</a></li>
        <li>部門情報追加</li>
    </ul>
</nav>

{{ if .Flg}}
<section id="errorMsg">
    <p>以下のメッセージをご確認ください。</p>

{{ range .ValidationMsg}}
<ul>
    <li>{{ .ValidationMsg }}</li>
</ul>
{{end}}

</section>
{{ end }}

<section>
    <p>情報を入力し、登録ボタンをクリックしてください</p>

    <form action=" /add/now" method="post">
        <table>
            <tbody>
            <tr>
                <th>部門番号&nbsp;<span class="required">必須</span></th>
                <td><input type="text" name="addDeptDeptno" id="addDeptDeptno" value=""></td>
            </tr>
            <tr>
                <th>部門名&nbsp;<span class="required">必須</span></th>
                <td><input type="text" name="addDeptDname" id="addDeptDname" value="" ></td>
            </tr>
            <tr>
                <th>所在地</th>
                <td><input type="text" name="addDeptLoc" id="addDeptLoc" value=""></td>
            </tr>
            <tr>
                <td colspan="2" class="submit"><input type="submit" value="登録"></td>
            </tr>
            </tbody>
        </table>
    </form>
</section>
</body>
</html>