<!DOCTYPE html>
<html lang="ja">
<head>
    <title>部門情報編集｜ScottAdmin</title>
    <link rel="stylesheet" type="text/css" href=" ../css/main.css">
    <meta charset="utf-8">
</head>
<body>
<h1>部門情報編集</h1>
<nav>
    <ul>
        <li><a href=" /">TOP</a></li>
        <li><a href=" /view">部門リスト</a></li>
        <li>部門情報編集</li>
    </ul>
</nav>
<section>
    <p>情報を入力し、更新ボタンをクリックしてください</p>

    <form action=" /edit/now" method="post">
        <table>
            <tbody>
            <tr>
                <th>部門番号</th>
                <td>{{ .Deptno }}<input type="hidden" name="editDeptDeptno" id="editDeptDeptno" value="{{ .Deptno}}"></td>
            </tr>
            <tr>
                <th>部門名&nbsp;<span class="required">必須</span></th>
                <td><input type="text" name="editDeptDname" id="editDeptDname" value="{{ .Dname}}" required></td>
            </tr>
            <tr>
                <th>所在地</th>
                <td><input type="text" name="editDeptLoc" id="editDeptLoc" value="{{ .Loc }}"></td>
            </tr>
            <tr>
                <td colspan="2" class="submit"><input type="submit" value="更新"></td>
            </tr>
            </tbody>
        </table>
    </form>
</section>
</body>
</html>