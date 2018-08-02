package controllers;

import entity.Museum;
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

public class MuseumListControl extends ControlWithSide implements Initializable {
    public static Museum[] museums;
    @FXML
    Button addMuseumToTrip;
    @FXML
    private TableView museumTable;
    @FXML
    private TableColumn museumID;
    @FXML
    private TableColumn museumName;
    @FXML
    private TableColumn museumLocation;
    @FXML
    private TableColumn museumPrice;
    @FXML
    private TableColumn museumOpenedAt;
    @FXML
    private TableColumn museumClosedAt;
    @FXML
    private TableColumn museumType;
    @FXML
    private TableColumn museumInfo;


    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setMuseumsTable();
    }

    private void setMuseumsTable() {
        ObservableList museumsList = FXCollections.observableList(Arrays.asList(museums));
        museumTable.setItems(museumsList);
        museumID.setCellValueFactory(new PropertyValueFactory("ID"));
        museumName.setCellValueFactory(new PropertyValueFactory("Name"));
        museumLocation.setCellValueFactory(new PropertyValueFactory("Location"));
        museumPrice.setCellValueFactory(new PropertyValueFactory("Price"));
        museumOpenedAt.setCellValueFactory(new PropertyValueFactory("OpenedAt"));
        museumClosedAt.setCellValueFactory(new PropertyValueFactory("ClosedAt"));
        museumType.setCellValueFactory(new PropertyValueFactory("MuseumType"));
        museumInfo.setCellValueFactory(new PropertyValueFactory("Info"));
        museumTable.getColumns().setAll(museumID, museumName, museumLocation, museumPrice, museumOpenedAt,museumClosedAt,museumType,museumInfo);
    }
}
