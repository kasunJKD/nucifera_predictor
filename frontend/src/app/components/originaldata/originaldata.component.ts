import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-originaldata',
  templateUrl: './originaldata.component.html',
  styleUrls: ['./originaldata.component.sass']
})
export class OriginaldataComponent implements OnInit  {
  data: any;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    // Send an HTTP POST request to the API endpoint when the component is initialized.
    const batchId = 1; // Set your batch ID here
    const requestData = { batchId : batchId };

    this.http.post('http://localhost:11001/api/batch/original/data/get', requestData).subscribe((response) => {
      this.data = response;
      console.log(response);
    });
  }
}
