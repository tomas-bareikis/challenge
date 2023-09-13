#!/usr/bin/awk -f
{
    s = 0;
    noZero = 1;
    for(i=3; i<=NF; i++) {
        if ($i == 0) {
            noZero = 0;
        }
        s += $i * $2;
    }

    if (noZero == 1) {
        found = 0
        k = 0
        for (i in names) {
            if (names[i] == $1) {
                all[i] += s;
                noZeros[i] += noZero;
                found = 1;
            }
            k++
        }

        if (found == 0) {
            names[k] = $1
            all[k] = s;
            noZeros[k] = noZero;
        }
    }
}
END{
    k = 0
    for (i in names) { k++ }
    for (i=0 ; i<k ; i++) {
        printf "%s %d %.2f\n", names[i], noZeros[i], all[i]/100000;
    }
}