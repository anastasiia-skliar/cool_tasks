package rest;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import entity.Event;
import entity.User;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;

public class UserService extends Headers {

    private static final String REST_SERVICE_URI = "http://cool-tasks.herokuapp.com/v1/";

    /*
     * Send a POST request to login in system.
     */
    public User Login(String login, String password) {
        RestTemplate restTemplate = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        java.net.CookieManager manager = new java.net.CookieManager();
        java.net.CookieHandler.setDefault(manager);
        User []user = null;
        try {
           String id = restTemplate.exchange(REST_SERVICE_URI + "login?login={login}&password={password}", HttpMethod.POST, request, String.class, login, password).getBody();
            user = getUserByID(id);
            return user[0];
        } catch (RuntimeException x) {
            return null;
        }
    }


    /*
     * Send a POST request to logout from the system.
     */
    public void Logout() {
        RestTemplate restTemplate = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        try {
            restTemplate.exchange(REST_SERVICE_URI + "logout", HttpMethod.POST, request, String.class).getBody();
        } catch (RuntimeException x) {
           System.exit(1);
        }
    }



    /*
     * Send a GET request to get a specific user.
     */
    public User[] getUserByID(String id) {
        RestTemplate restTemplate = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        String str  = restTemplate.exchange(REST_SERVICE_URI+"users?id="+id, HttpMethod.GET, request, String.class).getBody();
        Gson gson = new Gson();
        User user[] = gson.fromJson(str, User[].class);
        return user;
    }



}
