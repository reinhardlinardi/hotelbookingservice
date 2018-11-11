package helper;

import java.sql.*;

public class DBHandler {
    /* attributes */
    String username;
    String pass;
    String url = "jdbc:mysql://localhost:3306/hotel_booking";
    final int ERROR = -999;

    /* constructor */
    public DBHandler(){}

    public DBHandler(String username, String pass){
        this.username = username;
        this.pass = pass;
    }

    /* methods */
    public void executeQueryProcedure(String query){
        /* generic query executor without expecting any return */
        try{
            Class.forName("com.mysql.jdbc.Driver");
            Connection conn = DriverManager.getConnection(this.url,this.username,this.pass);

            ResultSet resultSet;

            PreparedStatement preparedStatement= conn.prepareStatement(query);

            resultSet = preparedStatement.executeQuery();
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    public int getCustomerId(String name, String email, String identity){
        int result = ERROR;
        try{
            Class.forName("com.mysql.jdbc.Driver");

            Connection conn = DriverManager.getConnection(this.url,this.username,this.pass);

            ResultSet resultSet;

            PreparedStatement preparedStatement= conn.prepareStatement("SELECT customer.id FROM customer WHERE (customer.name = ? AND customer.email = ? AND customer.identity = ?);");
            preparedStatement.setString(1, name);
            preparedStatement.setString(2, email);
            preparedStatement.setString(3, identity);

            resultSet = preparedStatement.executeQuery();

            if(resultSet.next()) {
                result = resultSet.getInt("id");
            }

            System.out.println(result);

        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        } catch (SQLException e) {
            e.printStackTrace();
        }
        return result;
    }
}
