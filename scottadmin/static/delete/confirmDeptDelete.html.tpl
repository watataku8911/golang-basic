<!DOCTYPE html>
<html lang="ja">
<head>
    <title>部門情報削除｜ScottAdmin</title>
    <link rel="stylesheet" type="text/css" href=" ../css/main.css">
    <meta charset="utf-8">
</head>
<body>
<h1>部門情報削除</h1>
<nav>
    <ul>
        <li><a href=" /">TOP</a></li>
        <li><a href=" /view">部門リスト</a></li>
        <li>部門情報削除</li>
    </ul>
</nav>

<section>
    <p>以下の部門情報を削除します。よろしければ、削除ボタンを押してください。</p>

    <form action=" /delete/now " method="post">
        <table border="1">
            <tbody>
            <tr>
                <th>部門番号</th>
                <td>{{ .Deptno }}<input type="hidden" name="deleteDeptDeptno" id="deleteDeptDeptno" value="{{ .Deptno}}"></td>
            </tr>
            <tr>
                <th>部門名</th>
                <td>{{ .Dname }}</td>
            </tr>
            <tr>
                <th>所在地</th>
                <td>{{ .Loc }}</td>
            </tr>
            <tr>
                <td colspan="2" class="submit"><input type="submit" value="削除"></td>
            </tr>
            </tbody>
        </table>
    </form>
</section>
</body>
</html>