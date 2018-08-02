package controllers;

import entity.User;
import javafx.event.ActionEvent;
import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.scene.Node;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.PasswordField;
import javafx.scene.control.TextField;
import javafx.stage.Stage;
import rest.UserService;
import util.AlertDialog;

import java.io.IOException;

public class LoginControl {

    @FXML
    TextField userLogin;
    @FXML
    PasswordField userPassword;

    private UserService userService = new UserService();

    public static User userInfo;

    public void logIn(ActionEvent add) throws IOException {
        if (userLogin.getText() != null && userPassword.getText() != null) {
            User user = null;
            user = userService.Login(userLogin.getText(),userPassword.getText());
            if (user != null) {
                userInfo = user;
                FXMLLoader fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/profile.fxml"));
                Parent root = (Parent) fxmlLoader.load();
                Stage stage = new Stage();
                stage.setScene(new Scene(root));
                stage.show();
                ((Node) (add.getSource())).getScene().getWindow().hide();
            } else {
                AlertDialog.display("Login error", "Something went wrong..");
                FXMLLoader fxmlLoader = new FXMLLoader(getClass().getClassLoader().getResource("Views/login.fxml"));
                Parent root = (Parent) fxmlLoader.load();
                Stage stage = new Stage();
                stage.setScene(new Scene(root));
                stage.show();
                ((Node) (add.getSource())).getScene().getWindow().hide();
            }
        }
    }
}