package controllers;

import entity.Restaurant;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Button;
import javafx.scene.control.TableColumn;
import javafx.scene.control.TableView;
import javafx.scene.control.cell.PropertyValueFactory;

import java.net.URL;
import java.util.Arrays;
import java.util.ResourceBundle;

public class RestaurantListControl extends ControlWithSide implements Initializable {

    public static Restaurant[] restaurants;
    @FXML
    Button addRestaurantToTrip;
    @FXML
    private TableView restaurantTable;
    @FXML
    private TableColumn restaurantID;
    @FXML
    private TableColumn restaurantName;
    @FXML
    private TableColumn restaurantLocation;
    @FXML
    private TableColumn restaurantStars;
    @FXML
    private TableColumn restaurantPrices;
    @FXML
    private TableColumn restaurantDescription;

    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setEventsRestaurant();
    }
    private void setEventsRestaurant() {
        ObservableList eventsList = FXCollections.observableList(Arrays.asList(restaurants));
        restaurantTable.setItems(eventsList);
        restaurantID.setCellValueFactory(new PropertyValueFactory("ID"));
        restaurantName.setCellValueFactory(new PropertyValueFactory("Name"));
        restaurantLocation.setCellValueFactory(new PropertyValueFactory("Location"));
        restaurantStars.setCellValueFactory(new PropertyValueFactory("Stars"));
        restaurantPrices.setCellValueFactory(new PropertyValueFactory("Prices"));
        restaurantDescription.setCellValueFactory(new PropertyValueFactory("Description"));
        restaurantTable.getColumns().setAll(restaurantID, restaurantName, restaurantLocation, restaurantStars, restaurantPrices, restaurantDescription);
    }
}
