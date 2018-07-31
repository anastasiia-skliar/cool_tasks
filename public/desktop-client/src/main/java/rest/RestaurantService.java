package rest;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import entity.Restaurant;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;

public class RestaurantService extends Headers {
    private static final String REST_SERVICE_URI = "http://cool-tasks.herokuapp.com/v1/";

    /*
     * Send a GET request to query for all Restaurants available.
     */

    public Restaurant[] getAllRestaurants() {
        RestTemplate rest = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        String jsonStr = rest.exchange(REST_SERVICE_URI + "restaurants", HttpMethod.GET, request, String.class).getBody();
        Gson gson = new GsonBuilder().setDateFormat("EEE, dd MMM yyyy HH:mm:ss zzz").create();
        return gson.fromJson(jsonStr, Restaurant[].class);
    }
}
