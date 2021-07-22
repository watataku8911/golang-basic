
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="utf-8">
    <title>部門情報リスト｜ScottAdmin</title>
    <link rel="stylesheet" type="text/css" href=" ../css/main.css">
</head>
<body>
<h1>部門情報リスト</h1>
<nav>
    <ul>
        <li><a href="/">TOP</a></li>
        <li>部門情報リスト</li>
    </ul>
</nav>
{{/*{{ template "flash" }}*/}}
<section>
    <p>
        新規登録は<a href="/add">こちら</a>から
    </p>
</section>


<section>
    <table>
        <thead>
        <tr>
            <th>部門番号</th>
            <th>部門名</th>
            <th>所在地</th>
            <th colspan="2">操作</th>
        </tr>
        </thead>
        <tbody>
        {{ range .}}
        <tr>
            <td>{{.Deptno}}</td>
            <td>{{.Dname}}</td>
            <td>{{.Loc}}</td>
            <td>
                <form action=" /edit" method="post">
                    <input type="hidden" name="editDeptDeptno" id="editDeptDeptno" value="{{.Deptno}}">
                    <input type="submit" value="編集">
                </form>
            </td>
            <td>
                <form action=" /delete" method="post">
                    <input type="hidden" name="deleteDeptDeptno" id="deleteDeptDeptno" value="{{.Deptno}}">
                    <input type="submit" value="削除">
                </form>
            </td>
        </tr>
        {{ end }}
        </tbody>
    </table>
</section>

</body>
</html>
