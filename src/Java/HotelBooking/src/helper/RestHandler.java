package helper;

import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

public class RestHandler {
    public RestHandler(){}

    /*
    input:
    url = url
    method = GET, POST, etc
    output: JSONObject
     */
    public JSONObject getRestObject(String in_url, String method){
        JSONObject jsonObject = null;

        try{
            URL url = new URL(in_url);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod(method);
            conn.setRequestProperty("Accept","application/json");

            if(conn.getResponseCode() != 200){
                throw new RuntimeException("Failed : HTTP error code : "+conn.getResponseCode());
            }

            BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));

            String output;


            System.out.println("Output from server...\n");
            while ((output = br.readLine())!= null){
                System.out.println(output);
                jsonObject = new JSONObject(output);
            }

            conn.disconnect();
        } catch(MalformedURLException e){
            e.printStackTrace();
        } catch(IOException e) {
            e.printStackTrace();
        }

        return(jsonObject);
    }

    public void deleteRESTRequest(String in_url){
        try{
            URL url = new URL(in_url);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("DELETE");
            conn.setRequestProperty("Accept","application/json");

            if(conn.getResponseCode() != 200){
                throw new RuntimeException("Failed : HTTP error code : "+conn.getResponseCode());
            }

            BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));

            String output;


            System.out.println("Output from server...\n");
            while ((output = br.readLine())!= null){
                System.out.println(output);
            }

            conn.disconnect();
        } catch(MalformedURLException e){
            e.printStackTrace();
        } catch(IOException e) {
            e.printStackTrace();
        }
    }
}
