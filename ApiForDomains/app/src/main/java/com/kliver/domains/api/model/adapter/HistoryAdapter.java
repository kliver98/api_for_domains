package com.kliver.domains.api.model.adapter;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.kliver.domains.api.R;

import java.util.ArrayList;

public class HistoryAdapter extends RecyclerView.Adapter<HistoryAdapter.ViewHolderDatos> {

    private ArrayList<String> items;
    private OnItemClickListener listener;

    public HistoryAdapter(ArrayList<String> items) {
        this.items = items;
    }

    public interface OnItemClickListener{
        void onItemClick(int position);
    }

    public void setOnItemClickListener(OnItemClickListener listener) {
        this.listener = listener;
    }

    @NonNull
    @Override
    public ViewHolderDatos onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_history,parent,false);
        return new ViewHolderDatos(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolderDatos holder, int position) {
        holder.addItem(items.get(position));
    }

    @Override
    public int getItemCount() {
        return items.size();
    }

    public class ViewHolderDatos extends RecyclerView.ViewHolder {

        private TextView item;

        public ViewHolderDatos(@NonNull View itemView) {
            super(itemView);
            item = itemView.findViewById(R.id.itemTV);

            itemView.setOnClickListener(new View.OnClickListener() {
                @Override
                public void onClick(View v) {
                    if (listener !=null) {
                        int position = getAdapterPosition();
                        if (position!=RecyclerView.NO_POSITION)
                            listener.onItemClick(position);
                    }
                }
            });
        }

        public void addItem(String dato) {
            this.item.setText(dato);
        }

    }

}
