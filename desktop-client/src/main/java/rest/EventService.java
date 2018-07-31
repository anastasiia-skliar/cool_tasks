package rest;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import entity.Event;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;

public class EventService extends Headers {
    private static final String REST_SERVICE_URI = "http://cool-tasks.herokuapp.com/v1/";

    /*
     * Send a GET request to query for all Events available.
     */

    public Event[] getAllEvents() {
        RestTemplate rest = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        String jsonStr = rest.exchange(REST_SERVICE_URI + "events", HttpMethod.GET, request, String.class).getBody();
        Gson gson = new GsonBuilder().setDateFormat("EEE, dd MMM yyyy HH:mm:ss zzz").create();
        return gson.fromJson(jsonStr, Event[].class);
    }
}
