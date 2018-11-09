package hotelbooking;

import java.sql.*;

public class Mainan {
    static final String JDBC_DRIVER = "com.mysql.jdbc.Driver";
    static final String DB_URL = "jdbc:mysql://localhost:3306/hotel_booking";

    static final String USER = "root";
    static final String PASS = "";

    public static void main(String[] args) {
        Connection conn = null;
        String name = "Cisco";
        String address = "di hatimuu";
        String job_desc = "gorengan";


        try {
            Class.forName(JDBC_DRIVER);

            System.out.println("\n Connecting to database");
            conn = DriverManager.getConnection(DB_URL,USER,PASS);
            System.out.println("\n SUCCESS");

            System.out.println("\n Inserting records into table");

            String sql = "INSERT INTO employee (name,address,job_desc)"+"VALUES(?,?,?)";

            PreparedStatement preparedStatement = conn.prepareStatement(sql);
            preparedStatement.setString(1,name);
            preparedStatement.setString(2,address);
            preparedStatement.setString(3,job_desc);
            preparedStatement.executeUpdate();

            System.out.println("SUCCESS!");
        } catch (SQLException se){
            se.printStackTrace();
        } catch(Exception e){
            e.printStackTrace();
        }

        System.out.println("FINALLY DONE YEAY");
    }
}
