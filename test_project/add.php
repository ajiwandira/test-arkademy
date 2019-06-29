<?php
// Create database connection using config file
include_once("config.php");
?>
<html>
<head>
    <title>Add Users</title>
</head>

<body>
    <a href="index.php">Go to Home</a>
    <br/><br/>

    <form action="add.php" method="post" name="form1">
        <table width="25%" border="0">
            <tr> 
                <td>Name</td>
                <td><input type="text" name="name"></td>
            </tr>
            <tr> 
                <td>Hobby</td>
                <td><select name="hobby">
                <?php $result11 = mysqli_query($mysqli, "SELECT * FROM hobi ORDER BY id");
        while($user_data11 = mysqli_fetch_array($result11)) {         
            ?>
  <option value="<?php echo $user_data11['id'] ?>"><?php echo $user_data11['name'] ?></option>
  <?php      
}
  ?>
</select></td>
            </tr>
            <tr> 
                <td>Category</td>
                <td><select name="category">
                <?php $result12 = mysqli_query($mysqli, "SELECT * FROM kategori ORDER BY id");
        while($user_data12 = mysqli_fetch_array($result12)) {         
            ?>
  <option value="<?php echo $user_data12['id'] ?>"><?php echo $user_data12['name'] ?></option>
  <?php        
}
  ?>
</select></td>
            </tr>
            <tr> 
                <td></td>
                <td><input type="submit" name="Submit" value="Add"></td>
            </tr>
        </table>
    </form>

    <?php

    // Check If form submitted, insert form data into users table.
    if(isset($_POST['Submit'])) {
        $name = $_POST['name'];
        $hobby = $_POST['hobby'];
        $category = $_POST['category'];

        // include database connection file
        include_once("config.php");
        // Insert user data into table

        $result5 = mysqli_query($mysqli, "INSERT INTO nama(name,id_hobby, id_category) VALUES('$name','$hobby','$category')");

        // Show message when user added
        echo "User added successfully. <a href='index.php'>View Users</a>";
    }
    ?>
</body>
</html>