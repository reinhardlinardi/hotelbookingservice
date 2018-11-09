package hotelbooking;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

public class RestClient {
    public static void main(String[] argv){

        try{
            URL url = new URL("http://localhost:8060/room/1");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");
            conn.setRequestProperty("Accept","application/json");

            if(conn.getResponseCode() != 200){
                throw new RuntimeException("Failed : HTTP error code : "+conn.getResponseCode());
            }

            BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));

            String output;

            System.out.println("Output from server...\n");
            while ((output = br.readLine())!= null){
                System.out.println(output);
//                JSONObject jsonObject = new JSONObject(output);
            }

            conn.disconnect();
        } catch(MalformedURLException e){
            e.printStackTrace();
        } catch(IOException e) {
            e.printStackTrace();
        }
    }
}
