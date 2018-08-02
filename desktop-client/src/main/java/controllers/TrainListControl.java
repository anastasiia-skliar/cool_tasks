package controllers;

import entity.Train;
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

public class TrainListControl extends ControlWithSide implements Initializable {
    public static Train[] trains;

    @FXML
    Button addTrainToTrip;
    @FXML
    private TableView trainTable;
    @FXML
    private TableColumn trainID;
    @FXML
    private TableColumn departureTime;
    @FXML
    private TableColumn arrivalTime;
    @FXML
    private TableColumn departureCity;
    @FXML
    private TableColumn arrivalCity;
    @FXML
    private TableColumn trainType;
    @FXML
    private TableColumn carType;
    @FXML
    private TableColumn price;


    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setEventsTable();
    }
    private void setEventsTable() {
        ObservableList trainList = FXCollections.observableList(Arrays.asList(trains));
        trainTable.setItems(trainList);
        trainID.setCellValueFactory(new PropertyValueFactory("ID"));
        departureTime.setCellValueFactory(new PropertyValueFactory("DepartureTime"));
        arrivalTime.setCellValueFactory(new PropertyValueFactory("ArrivalTime"));
        departureCity.setCellValueFactory(new PropertyValueFactory("DepartureCity"));
        arrivalCity.setCellValueFactory(new PropertyValueFactory("ArrivalCity"));
        trainType.setCellValueFactory(new PropertyValueFactory("TrainType"));
        carType.setCellValueFactory(new PropertyValueFactory("CarType"));
        price.setCellValueFactory(new PropertyValueFactory("Price"));
        trainTable.getColumns().setAll(trainID, departureTime, arrivalTime, departureCity, arrivalCity, trainType,carType,price);

    }
}
