package com.kliver.domains.api.view;

import androidx.appcompat.app.AppCompatActivity;

import com.kliver.domains.api.R;
import com.kliver.domains.api.control.HistoryController;

import android.os.Bundle;
import android.util.Log;
import android.widget.ArrayAdapter;
import android.widget.ImageView;
import android.widget.ListView;

import java.util.ArrayList;

public class HistoryActivity extends AppCompatActivity {

    private ImageView backArrowIV;
    private ListView listView;
    private ArrayAdapter<String> adapter;
    private ArrayList<String> items;
    private HistoryController controller;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_history);
        backArrowIV = findViewById(R.id.backArrowIV);
        listView = findViewById(R.id.listView);
        items = new ArrayList<String>();
        adapter = new ArrayAdapter<String>(this, android.R.layout.simple_list_item_1,items);
        listView.setAdapter(adapter);
        controller = new HistoryController(this);
    }

    public ImageView getBackArrowIV() {
        return backArrowIV;
    }

    public ListView getListView() {
        return listView;
    }

    public ArrayAdapter<String> getAdapter() {
        return adapter;
    }

    public ArrayList<String> getItems() {
        return items;
    }

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        this.finish();
    }

}
