package controllers;

import com.jfoenix.controls.JFXButton;
import javafx.event.ActionEvent;
import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.scene.Node;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.stage.Stage;
import rest.*;

import java.io.IOException;

public class SidePanelControl {

    private UserService userService = new UserService();
    private EventService eventService = new EventService();
    private HotelService hotelService = new HotelService();
    private RestaurantService restaurantService = new RestaurantService();
    private MuseumService museumService = new MuseumService();
    private FlightService flightService = new FlightService();
    private TrainService trainService = new TrainService();
    private TaskService taskService = new TaskService();

    @FXML
    private void changePage(ActionEvent event) throws IOException {

        JFXButton btn = (JFXButton) event.getSource();
        FXMLLoader fxmlLoader = null;
        switch(btn.getText())
        {
            case "Tasks":
                TaskListControl.tasks = taskService.getAllTasksForUser(LoginControl.userInfo.getID());
               fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/tasks.fxml"));
                break;
           case "Trips":
               break;
               case "Events":
                 EventListControl.events = eventService.getAllEvents();
                fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/events.fxml"));
                     break;
            case "Flights":
               FlightListControl.flights = flightService.getAllFlights();
                fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/flights.fxml"));
               break;
            case "Hotels":
                 HotelListControl.hotels = hotelService.getAllHotels();
                fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/hotels.fxml"));
                break;
            case "Museums":
                    MuseumListControl.museums = museumService.getAllMuseums();
                fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/museums.fxml"));
                break;
            case "Restaurants":
               RestaurantListControl.restaurants = restaurantService.getAllRestaurants();
                fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/restaurants.fxml"));
                break;
            case "Trains":
                TrainListControl.trains = trainService.getAllTrains();
                fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/trains.fxml"));
                break;
        }
        Parent root = (Parent) fxmlLoader.load();
        Stage stage = new Stage();
        stage.setScene(new Scene(root));
        stage.show();
        ((Node)(event.getSource())).getScene().getWindow().hide();
    }
    @FXML
    private void exit(ActionEvent event) {
        userService.Logout();
        System.exit(0);
    }
}
