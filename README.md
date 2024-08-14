# trimmunk
Genomic reads trimming project implemented in the Go-language. This is just an excercise, I'm not planning to take down trimmomatic, although someone should because I hate Java. I'm trying to learn Go more than I'm trying to implement something new but if someone likes the idea, that's cool. Ok, so this is the plan of action:

## Plan of Action
- [ ] **Reading .fastq files**
- [x] Trimmunk should put the sequences in single lines with the corresponding phred score under
- [ ] Trimmunk should save this sequences in a new, temporal file
- [ ] **Running statistics in .fastq sequences**
- [ ] Trimmunk should interpret the Phred scores based on ASCII characters.
- [ ] Trimmunk should do a list? So that you can easily read each character and its corresponding punctuation.
- [ ] Trimmunk should plot all the reads (maybe I'll use Python here)
- [ ] **Using functions to cut bad quality data**
- [ ] Trimmunk should filter (This should be easy peasy because this should be based in user input)
- [ ] Trimmunk should then calculate the slope and, based on this, give a report on where it will perform the cut.
