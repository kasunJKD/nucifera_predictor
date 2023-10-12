import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
  imageSrc: SafeResourceUrl; // Variable to store the image data
  data: any; // Variable to store the response data

  constructor(private http: HttpClient, private sanitizer: DomSanitizer) {}

  ngOnInit() {
    // Send an HTTP POST request to the API endpoint when the component is initialized.
    const batchId = 1; // Set your batch ID here
    const requestData = { batchNumber : batchId };

    this.http.post('http://localhost:11001/api/batch/data/get', requestData).subscribe((response) => {
      this.data = response;
      const base64Image = this.data.batchResponse[0].plotFit;
      const decodedImageData = atob(base64Image);
      const imageUrl = 'data:image/png;base64,' + decodedImageData;

      console.log(imageUrl);

      this.imageSrc = this.sanitizer.bypassSecurityTrustResourceUrl(imageUrl);
    });
  }
}
