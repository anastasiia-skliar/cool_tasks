package controllers;

import entity.Task;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.TableColumn;
import javafx.scene.control.TableView;
import javafx.scene.control.cell.PropertyValueFactory;

import java.net.URL;
import java.util.Arrays;
import java.util.ResourceBundle;

public class TaskListControl extends ControlWithSide implements Initializable {
    public static Task[] tasks;

    @FXML
    private TableView taskTable;
    @FXML
    private TableColumn taskID;
    @FXML
    private TableColumn taskName;
    @FXML
    private TableColumn taskTime;
    @FXML
    private TableColumn createdAt;
    @FXML
    private TableColumn updatedAt;
    @FXML
    private TableColumn Description;

    @Override
    public void initialize(URL location, ResourceBundle resources) {
        initializeSide(location, resources);
        setTaskTable();
    }
    private void setTaskTable() {
        ObservableList taskList = FXCollections.observableList(Arrays.asList(tasks));
        taskTable.setItems(taskList);
        taskID.setCellValueFactory(new PropertyValueFactory("ID"));
        taskName.setCellValueFactory(new PropertyValueFactory("Name"));
        taskTime.setCellValueFactory(new PropertyValueFactory("Time"));
        createdAt.setCellValueFactory(new PropertyValueFactory("CreatedAt"));
        updatedAt.setCellValueFactory(new PropertyValueFactory("UpdatedAt"));
        Description.setCellValueFactory(new PropertyValueFactory("Description"));
        taskTable.getColumns().setAll(taskID, taskName, taskTime, createdAt, updatedAt, Description);
    }

}
