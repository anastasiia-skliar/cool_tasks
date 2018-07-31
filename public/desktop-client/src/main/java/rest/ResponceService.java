package rest;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.util.JSONPObject;
import com.google.gson.Gson;
import entity.Hotel;
import entity.User;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.web.client.RestTemplate;

import java.util.Collection;
import java.util.List;

public class ResponceService extends Headers {

  public static final String REST_SERVICE_URI = "http://cool-tasks.herokuapp.com/v1/";

  public String getResponce() {
    return "";
  }

  public String method(String login, String password) {
    RestTemplate rest = new RestTemplate();
    HttpEntity<String> request = new HttpEntity<String>(getHeaders());
    java.net.CookieManager manager = new java.net.CookieManager();
    java.net.CookieHandler.setDefault(manager);
    String result2 = rest.exchange("http://localhost:8080/v1/login?login={login}&password={password}", HttpMethod.POST, request, String.class,login,password).getBody();
    System.out.println(result2);
    return result2;
  }

  public Hotel[] method2() {
    RestTemplate rest = new RestTemplate();
    HttpEntity<String> request = new HttpEntity<String>(getHeaders());
    String jsonStr = rest.exchange("http://localhost:8080/v1/hotels",HttpMethod.GET,request, String.class).getBody();
    Gson gson = new Gson();
    Hotel users[] = gson.fromJson(jsonStr, Hotel[].class);
    return  users;
  }
}
