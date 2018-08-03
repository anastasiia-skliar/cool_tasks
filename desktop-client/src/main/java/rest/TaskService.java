package rest;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import entity.Task;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;
import util.GsonUTCDateAdapter;

import java.util.Date;

public class TaskService extends Headers {

    private static final String REST_SERVICE_URI = "http://cool-tasks.herokuapp.com/v1/";

    /*
     * Send a GET request to query for all Events available.
     */
    public Task[] getAllTasksForUser(String id){
        RestTemplate rest = new RestTemplate();
        HttpEntity<String> request = new HttpEntity<String>(getHeaders());
        String jsonStr = rest.exchange(REST_SERVICE_URI+"tasks?=id"+id, HttpMethod.GET,request, String.class).getBody();
        Gson gson = new GsonBuilder().setDateFormat("yyyy-MM-dd'T'HH:mm:ssz").create();
        return gson.fromJson(jsonStr, Task[].class);
    }
}
