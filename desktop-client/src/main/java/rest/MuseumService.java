package rest;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import entity.Museum;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;
import util.GsonUTCDateAdapter;

import java.util.Date;

public class MuseumService extends Headers {
    private static final String REST_SERVICE_URI = "http://cool-tasks.herokuapp.com/v1/";
    /*
     * Send a GET request to query for all Museums available.
     */

    public Museum[] getAllMuseums() {
        RestTemplate rest = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        String jsonStr = rest.exchange(REST_SERVICE_URI+"museums", HttpMethod.GET, request, String.class).getBody();
        Gson gson = new GsonBuilder().registerTypeAdapter(Date.class, new GsonUTCDateAdapter()).create();
        return gson.fromJson(jsonStr, Museum[].class);
    }
}
