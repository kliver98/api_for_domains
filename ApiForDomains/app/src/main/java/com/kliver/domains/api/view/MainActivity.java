package com.kliver.domains.api.view;

import androidx.appcompat.app.AppCompatActivity;
import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.TextView;

import com.kliver.domains.api.R;
import com.kliver.domains.api.control.MainController;

public class MainActivity extends AppCompatActivity {

    private MainController controller;
    private Button btnDomain;
    private Button btnHistory;
    private TextView versionTV;
    private TextView serverTV;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        btnDomain = findViewById(R.id.btnDomain);
        btnHistory = findViewById(R.id.btnHistory);
        versionTV = findViewById(R.id.versionTV);
        serverTV = findViewById(R.id.serverTV);

        controller = new MainController(this);
    }

    public void openDomainActivity() {
        Intent i = new Intent(this,DomainActivity.class);
        this.startActivity(i);
    }

    public void openHistoryActivity() {
        Intent i = new Intent(this,HistoryActivity.class);
        this.startActivity(i);
    }

    public Button getBtnDomain() {
        return btnDomain;
    }

    public Button getBtnHistory() {
        return btnHistory;
    }

    public TextView getVersionTV() {
        return versionTV;
    }

    public TextView getServerTV() {
        return serverTV;
    }
}
